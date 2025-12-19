
import com.azure.resourcemanager.avs.models.AddonVrProperties;

/**
 * Samples for Addons CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Addons_CreateOrUpdate_VR.json
     */
    /**
     * Sample code: Addons_CreateOrUpdate_VR.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsCreateOrUpdateVR(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().define("vr").withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new AddonVrProperties().withVrsCount(1)).create();
    }
}
