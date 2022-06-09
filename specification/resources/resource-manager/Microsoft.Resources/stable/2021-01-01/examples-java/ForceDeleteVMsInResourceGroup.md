```java
import com.azure.core.util.Context;

/** Samples for ResourceGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/ForceDeleteVMsInResourceGroup.json
     */
    /**
     * Sample code: Force delete all the Virtual Machines in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void forceDeleteAllTheVirtualMachinesInAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getResourceGroups()
            .delete("my-resource-group", "Microsoft.Compute/virtualMachines", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
