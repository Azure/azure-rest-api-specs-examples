Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualMachineRunCommands GetByVirtualMachine. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/GetRunCommand.json
     */
    /**
     * Sample code: Get a run command.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineRunCommands()
            .getByVirtualMachineWithResponse("myResourceGroup", "myVM", "myRunCommand", null, Context.NONE);
    }
}
```
