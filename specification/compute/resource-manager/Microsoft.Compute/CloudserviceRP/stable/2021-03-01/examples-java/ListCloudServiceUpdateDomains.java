import com.azure.core.util.Context;

/** Samples for CloudServicesUpdateDomain ListUpdateDomains. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServiceUpdateDomains.json
     */
    /**
     * Sample code: List Update Domains in Cloud Service.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listUpdateDomainsInCloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServicesUpdateDomains()
            .listUpdateDomains("ConstosoRG", "{cs-name}", Context.NONE);
    }
}
