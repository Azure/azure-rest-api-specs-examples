
/**
 * Samples for AssociationsInterface ListByTrafficController.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/AssociationsGet.json
     */
    /**
     * Sample code: Get Associations.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getAssociations(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.associationsInterfaces().listByTrafficController("rg1", "tc1", com.azure.core.util.Context.NONE);
    }
}
