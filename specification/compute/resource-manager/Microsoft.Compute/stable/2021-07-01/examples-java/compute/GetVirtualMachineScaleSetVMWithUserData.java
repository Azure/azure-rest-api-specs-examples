import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/GetVirtualMachineScaleSetVMWithUserData.json
     */
    /**
     * Sample code: Get VM scale set VM with UserData.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVMScaleSetVMWithUserData(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .getWithResponse("myResourceGroup", "{vmss-name}", "0", null, Context.NONE);
    }
}
