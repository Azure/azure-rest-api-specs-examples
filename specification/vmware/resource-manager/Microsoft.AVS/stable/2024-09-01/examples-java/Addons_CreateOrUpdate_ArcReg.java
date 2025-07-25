
import com.azure.resourcemanager.avs.models.AddonArcProperties;

/**
 * Samples for Addons CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Addons_CreateOrUpdate_ArcReg.json
     */
    /**
     * Sample code: Addons_CreateOrUpdate_ArcReg.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsCreateOrUpdateArcReg(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().define("arc").withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new AddonArcProperties().withVCenter(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_test/providers/Microsoft.ConnectedVMwarevSphere/VCenters/test-vcenter"))
            .create();
    }
}
