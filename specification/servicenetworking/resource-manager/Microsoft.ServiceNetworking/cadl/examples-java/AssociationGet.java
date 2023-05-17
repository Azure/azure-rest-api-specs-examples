/** Samples for AssociationsInterface Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationGet.json
     */
    /**
     * Sample code: Get Association.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getAssociation(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.associationsInterfaces().getWithResponse("rg1", "tc1", "as1", com.azure.core.util.Context.NONE);
    }
}
