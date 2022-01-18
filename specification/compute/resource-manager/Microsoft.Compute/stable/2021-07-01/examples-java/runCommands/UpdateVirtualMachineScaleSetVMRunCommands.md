Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandScriptSource;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandUpdate;

/** Samples for VirtualMachineScaleSetVMRunCommands Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/UpdateVirtualMachineScaleSetVMRunCommands.json
     */
    /**
     * Sample code: Update VirtualMachineScaleSet VM run command.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVirtualMachineScaleSetVMRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMRunCommands()
            .update(
                "myResourceGroup",
                "myvmScaleSet",
                "0",
                "myRunCommand",
                new VirtualMachineRunCommandUpdate()
                    .withSource(
                        new VirtualMachineRunCommandScriptSource().withScript("Write-Host Script Source Updated!")),
                Context.NONE);
    }
}
```
