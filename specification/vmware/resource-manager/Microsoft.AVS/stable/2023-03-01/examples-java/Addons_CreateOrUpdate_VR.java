import com.azure.resourcemanager.avs.models.AddonVrProperties;

/** Samples for Addons CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Addons_CreateOrUpdate_VR.json
     */
    /**
     * Sample code: Addons_CreateOrUpdate_VR.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void addonsCreateOrUpdateVR(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .addons()
            .define("vr")
            .withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new AddonVrProperties().withVrsCount(1))
            .create();
    }
}
