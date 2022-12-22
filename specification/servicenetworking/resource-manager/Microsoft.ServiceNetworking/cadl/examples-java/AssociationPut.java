import com.azure.resourcemanager.servicenetworking.models.AssociationSubnet;
import com.azure.resourcemanager.servicenetworking.models.AssociationType;

/** Samples for AssociationsInterface CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationPut.json
     */
    /**
     * Sample code: Put Association.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void putAssociation(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager
            .associationsInterfaces()
            .define("associatedvnet-1")
            .withRegion("West US")
            .withExistingTrafficController("rg1", "TC1")
            .withAssociationType(AssociationType.SUBNETS)
            .withSubnet(new AssociationSubnet().withId("subnetFullRef"))
            .create();
    }
}
