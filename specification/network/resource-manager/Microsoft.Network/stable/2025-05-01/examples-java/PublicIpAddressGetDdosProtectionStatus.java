
/**
 * Samples for PublicIpAddresses DdosProtectionStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * PublicIpAddressGetDdosProtectionStatus.json
     */
    /**
     * Sample code: Get Ddos Protection Status of a Public IP Address.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDdosProtectionStatusOfAPublicIPAddress(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().ddosProtectionStatus("rg1", "test-pip",
            com.azure.core.util.Context.NONE);
    }
}
