
import com.azure.resourcemanager.vmwarecloudsimple.models.Sku;
import java.util.UUID;

/**
 * Samples for DedicatedCloudNodes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/
     * CreateDedicatedCloudNode.json
     */
    /**
     * Sample code: CreateDedicatedCloudNode.
     * 
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void
        createDedicatedCloudNode(com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.dedicatedCloudNodes().define("myNode").withRegion("westus").withExistingResourceGroup("myResourceGroup")
            .withSku(new Sku().withName("VMware_CloudSimple_CS28")).withAvailabilityZoneId("az1").withNodesCount(1)
            .withPlacementGroupId("n1").withPurchaseId(UUID.fromString("56acbd46-3d36-4bbf-9b08-57c30fdf6932"))
            .withIdPropertiesId("general").withNamePropertiesName("CS28-Node")
            .withReferer("https://management.azure.com/").create();
    }
}
