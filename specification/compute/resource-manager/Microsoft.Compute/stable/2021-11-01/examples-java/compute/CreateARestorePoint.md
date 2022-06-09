```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.RestorePointInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import java.util.Arrays;

/** Samples for RestorePoints Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateARestorePoint.json
     */
    /**
     * Sample code: Create a restore point.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createARestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
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
                    .withExcludeDisks(
                        Arrays
                            .asList(
                                new ApiEntityReference()
                                    .withId(
                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
