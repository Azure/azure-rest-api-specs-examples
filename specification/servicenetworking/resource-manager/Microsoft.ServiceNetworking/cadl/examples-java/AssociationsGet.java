import com.azure.core.util.Context;

/** Samples for AssociationsInterface ListByTrafficController. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationsGet.json
     */
    /**
     * Sample code: Get Associations.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getAssociations(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.associationsInterfaces().listByTrafficController("rg1", "TC1", Context.NONE);
    }
}
