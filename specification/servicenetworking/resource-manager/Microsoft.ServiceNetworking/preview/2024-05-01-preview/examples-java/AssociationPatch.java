
import com.azure.resourcemanager.servicenetworking.models.Association;
import com.azure.resourcemanager.servicenetworking.models.AssociationSubnetUpdate;
import com.azure.resourcemanager.servicenetworking.models.AssociationType;
import com.azure.resourcemanager.servicenetworking.models.AssociationUpdateProperties;

/**
 * Samples for AssociationsInterface Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/preview/2024-05-01-preview/examples/
     * AssociationPatch.json
     */
    /**
     * Sample code: Update Association.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void updateAssociation(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        Association resource = manager.associationsInterfaces()
            .getWithResponse("rg1", "tc1", "as1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new AssociationUpdateProperties().withAssociationType(AssociationType.SUBNETS)
            .withSubnet(new AssociationSubnetUpdate().withId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-tc/subnets/tc-subnet")))
            .apply();
    }
}
