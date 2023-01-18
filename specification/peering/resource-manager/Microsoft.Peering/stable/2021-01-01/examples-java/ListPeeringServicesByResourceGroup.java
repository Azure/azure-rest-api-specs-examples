/** Samples for PeeringServices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeeringServicesByResourceGroup.json
     */
    /**
     * Sample code: List peering services in a resource group.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeeringServicesInAResourceGroup(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peeringServices().listByResourceGroup("rgName", com.azure.core.util.Context.NONE);
    }
}
