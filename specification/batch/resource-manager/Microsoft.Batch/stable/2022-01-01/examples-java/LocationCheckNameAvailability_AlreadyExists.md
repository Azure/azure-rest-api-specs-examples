```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.CheckNameAvailabilityParameters;

/** Samples for Location CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/LocationCheckNameAvailability_AlreadyExists.json
     */
    /**
     * Sample code: LocationCheckNameAvailability_AlreadyExists.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void locationCheckNameAvailabilityAlreadyExists(
        com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .locations()
            .checkNameAvailabilityWithResponse(
                "japaneast", new CheckNameAvailabilityParameters().withName("existingaccountname"), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
