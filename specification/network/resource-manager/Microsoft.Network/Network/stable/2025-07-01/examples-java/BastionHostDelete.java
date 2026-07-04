
/**
 * Samples for BastionHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostDelete.json
     */
    /**
     * Sample code: Delete Bastion Host.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteBastionHost(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().delete("rg1", "bastionhosttenant", com.azure.core.util.Context.NONE);
    }
}
