
/**
 * Samples for BastionHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostDeveloperDelete.json
     */
    /**
     * Sample code: Delete Developer Bastion Host.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteDeveloperBastionHost(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().delete("rg2", "bastionhostdeveloper",
            com.azure.core.util.Context.NONE);
    }
}
