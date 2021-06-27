package caller

import (
	"context"
	"fmt"

	"cloud.google.com/go/logging"
	"google.golang.org/api/compute/v1"
	// "google.golang.org/api/option"
)

const projID = "PROJECT_ID"

// const quotaProjID = "QUOTA_PROJECT_ID"

const zone = "us-central1-a"

func Execute() error {
	ctx := context.Background()
	// clientOption := option.WithQuotaProject(quotaProjID)

	// LIST GOOGLE CLOUD ENGINE (GCE) INSTANCES
	computeService, err := compute.NewService(ctx)
	// computeService, err := compute.NewService(ctx, clientOption)
	if err != nil {
		return err
	}
	instanceListCall := computeService.Instances.List(projID, zone)
	instanceList, err := instanceListCall.Do()
	if err != nil {
		return err
	}
	for _, instance := range instanceList.Items {
		fmt.Println(instance.Id)
	}

	// WRITE A LOG ENTRY TO CLOUD LOGGING
	loggingClient, err := logging.NewClient(ctx, projID)
	// loggingClient, err := logging.NewClient(ctx, projID, clientOption)
	defer loggingClient.Close()
	if err != nil {
		return err
	}
	logger := loggingClient.Logger("debug")
	logger.Log(logging.Entry{Payload: "hello world"})

	return nil
}
