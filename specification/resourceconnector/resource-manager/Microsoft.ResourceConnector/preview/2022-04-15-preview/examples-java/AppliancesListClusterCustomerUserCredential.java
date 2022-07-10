import com.azure.core.util.Context;

/** Samples for Appliances ListClusterCustomerUserCredential. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesListClusterCustomerUserCredential.json
     */
    /**
     * Sample code: ListClusterCustomerUserCredentialAppliance.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void listClusterCustomerUserCredentialAppliance(
        com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager
            .appliances()
            .listClusterCustomerUserCredentialWithResponse("testresourcegroup", "appliance01", Context.NONE);
    }
}
