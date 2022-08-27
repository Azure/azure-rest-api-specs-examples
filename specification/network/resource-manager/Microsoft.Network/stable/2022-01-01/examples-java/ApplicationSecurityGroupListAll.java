import com.azure.core.util.Context;

/** Samples for ApplicationSecurityGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ApplicationSecurityGroupListAll.json
     */
    /**
     * Sample code: List all application security groups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllApplicationSecurityGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationSecurityGroups().list(Context.NONE);
    }
}
