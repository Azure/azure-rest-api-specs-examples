import com.azure.core.util.Context;

/** Samples for DscpConfiguration Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/DscpConfigurationDelete.json
     */
    /**
     * Sample code: Delete DSCP Configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteDSCPConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDscpConfigurations().delete("rg1", "mydscpConfig", Context.NONE);
    }
}
