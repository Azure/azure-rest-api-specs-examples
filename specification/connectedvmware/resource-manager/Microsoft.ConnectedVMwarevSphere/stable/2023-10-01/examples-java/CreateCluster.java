
import com.azure.resourcemanager.connectedvmware.models.ExtendedLocation;

/**
 * Samples for Clusters Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/
     * CreateCluster.json
     */
    /**
     * Sample code: CreateCluster.
     * 
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void createCluster(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().define("HRCluster").withRegion("East US").withExistingResourceGroup("testrg")
            .withExtendedLocation(new ExtendedLocation().withType("customLocation").withName(
                "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"))
            .withVCenterId(
                "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter")
            .withMoRefId("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee").create();
    }
}
