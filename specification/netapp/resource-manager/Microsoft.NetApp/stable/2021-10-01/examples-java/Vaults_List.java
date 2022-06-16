import com.azure.core.util.Context;

/** Samples for Vaults List. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Vaults_List.json
     */
    /**
     * Sample code: Vaults_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void vaultsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.vaults().list("myRG", "account1", Context.NONE);
    }
}
