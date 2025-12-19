
/**
 * Samples for PrivateClouds GetVcfLicense.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_GetVcfLicense.json
     */
    /**
     * Sample code: PrivateClouds_GetVcfLicense.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsGetVcfLicense(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().getVcfLicenseWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
