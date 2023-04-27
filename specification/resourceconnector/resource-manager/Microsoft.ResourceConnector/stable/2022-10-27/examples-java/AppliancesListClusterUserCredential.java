/** Samples for Appliances ListClusterUserCredential. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListClusterUserCredential.json
     */
    /**
     * Sample code: ListClusterUserCredentialAppliance.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void listClusterUserCredentialAppliance(
        com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager
            .appliances()
            .listClusterUserCredentialWithResponse(
                "testresourcegroup", "appliance01", com.azure.core.util.Context.NONE);
    }
}
