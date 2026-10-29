
import com.azure.resourcemanager.cdn.models.CanMigrateParameters;
import com.azure.resourcemanager.cdn.models.ResourceReference;

/**
 * Samples for Profiles CanMigrate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_CanMigrate.json
     */
    /**
     * Sample code: Profiles_CanMigrate.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesCanMigrate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().canMigrate("RG",
            new CanMigrateParameters().withClassicResourceReference(new ResourceReference()
                .withId("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoors/frontdoorname")),
            com.azure.core.util.Context.NONE);
    }
}
