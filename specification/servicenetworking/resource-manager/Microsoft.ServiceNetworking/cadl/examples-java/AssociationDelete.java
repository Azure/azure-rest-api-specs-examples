/** Samples for AssociationsInterface Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationDelete.json
     */
    /**
     * Sample code: Delete Association.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void deleteAssociation(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.associationsInterfaces().delete("rg1", "tc1", "as1", com.azure.core.util.Context.NONE);
    }
}
