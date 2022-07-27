import com.azure.core.util.Context;

/** Samples for CloudServicesUpdateDomain WalkUpdateDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudServiceUpdateDomain_Update.json
     */
    /**
     * Sample code: Update Cloud Service to specified Domain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateCloudServiceToSpecifiedDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServicesUpdateDomains()
            .walkUpdateDomain("ConstosoRG", "{cs-name}", 1, null, Context.NONE);
    }
}
