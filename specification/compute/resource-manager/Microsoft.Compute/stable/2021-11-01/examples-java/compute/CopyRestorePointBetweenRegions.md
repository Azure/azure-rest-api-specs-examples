```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.RestorePointInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;

/** Samples for RestorePoints Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CopyRestorePointBetweenRegions.json
     */
    /**
     * Sample code: Copy a restore point to a different region.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void copyARestorePointToADifferentRegion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePoints()
            .create(
                "myResourceGroup",
                "rpcName",
                "rpName",
                new RestorePointInner()
                    .withSourceRestorePoint(
                        new ApiEntityReference()
                            .withId(
                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName/restorePoints/sourceRpName")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
