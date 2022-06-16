import com.azure.core.util.Context;

/** Samples for RestorePoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetRestorePointWithInstanceView.json
     */
    /**
     * Sample code: Get restore point with instance view.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRestorePointWithInstanceView(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePoints()
            .getWithResponse("myResourceGroup", "rpcName", "rpName", null, Context.NONE);
    }
}
