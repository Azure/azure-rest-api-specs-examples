```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.ActivateApplicationPackageParameters;

/** Samples for ApplicationPackage Activate. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationPackageActivate.json
     */
    /**
     * Sample code: ApplicationPackageActivate.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationPackageActivate(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .applicationPackages()
            .activateWithResponse(
                "default-azurebatch-japaneast",
                "sampleacct",
                "app1",
                "1",
                new ActivateApplicationPackageParameters().withFormat("zip"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
