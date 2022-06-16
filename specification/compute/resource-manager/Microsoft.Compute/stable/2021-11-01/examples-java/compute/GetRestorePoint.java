import com.azure.core.util.Context;

/** Samples for RestorePoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetRestorePoint.json
     */
    /**
     * Sample code: Get a restore point.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePoints()
            .getWithResponse("myResourceGroup", "rpcName", "rpName", null, Context.NONE);
    }
}
