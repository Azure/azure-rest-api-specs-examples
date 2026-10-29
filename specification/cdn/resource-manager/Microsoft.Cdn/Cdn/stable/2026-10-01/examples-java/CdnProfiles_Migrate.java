
import com.azure.resourcemanager.cdn.models.CdnMigrationToAfdParameters;
import com.azure.resourcemanager.cdn.models.Sku;
import com.azure.resourcemanager.cdn.models.SkuName;

/**
 * Samples for Profiles CdnMigrateToAfd.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/CdnProfiles_Migrate.json
     */
    /**
     * Sample code: Profiles_CdnMigrate.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesCdnMigrate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().cdnMigrateToAfd("RG", "profile1",
            new CdnMigrationToAfdParameters().withSku(new Sku().withName(SkuName.STANDARD_AZURE_FRONT_DOOR)),
            com.azure.core.util.Context.NONE);
    }
}
