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
            .define("as1")
            .withRegion("NorthCentralUS")
            .withExistingTrafficController("rg1", "tc1")
            .withAssociationType(AssociationType.SUBNETS)
            .withSubnet(
                new AssociationSubnet()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-tc/subnets/tc-subnet"))
            .create();
    }
}
