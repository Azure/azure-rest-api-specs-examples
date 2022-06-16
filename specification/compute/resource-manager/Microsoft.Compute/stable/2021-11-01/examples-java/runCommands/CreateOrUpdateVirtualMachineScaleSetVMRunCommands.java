import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineRunCommandInner;
import com.azure.resourcemanager.compute.models.RunCommandInputParameter;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandScriptSource;
import java.util.Arrays;

/** Samples for VirtualMachineScaleSetVMRunCommands CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/CreateOrUpdateVirtualMachineScaleSetVMRunCommands.json
     */
    /**
     * Sample code: Create VirtualMachineScaleSet VM run command.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualMachineScaleSetVMRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMRunCommands()
            .createOrUpdate(
                "myResourceGroup",
                "myvmScaleSet",
                "0",
                "myRunCommand",
                new VirtualMachineRunCommandInner()
                    .withLocation("West US")
                    .withSource(new VirtualMachineRunCommandScriptSource().withScript("Write-Host Hello World!"))
                    .withParameters(
                        Arrays
                            .asList(
                                new RunCommandInputParameter().withName("param1").withValue("value1"),
                                new RunCommandInputParameter().withName("param2").withValue("value2")))
                    .withAsyncExecution(false)
                    .withRunAsUser("user1")
                    .withRunAsPassword("<runAsPassword>")
                    .withTimeoutInSeconds(3600),
                Context.NONE);
    }
}
