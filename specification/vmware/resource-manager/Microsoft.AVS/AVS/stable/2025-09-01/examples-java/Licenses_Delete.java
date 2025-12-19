
import com.azure.resourcemanager.avs.models.LicenseName;

/**
 * Samples for Licenses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Licenses_Delete.json
     */
    /**
     * Sample code: Licenses_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void licensesDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.licenses().delete("group1", "cloud1", LicenseName.VMWARE_FIREWALL, com.azure.core.util.Context.NONE);
    }
}
