
import com.azure.resourcemanager.appconfiguration.models.KeyValueFilter;
import java.util.Arrays;

/**
 * Samples for Snapshots Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateSnapshot.json
     */
    /**
     * Sample code: Snapshots_Create.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void snapshotsCreate(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.snapshots().define("mySnapshot").withExistingConfigurationStore("myResourceGroup", "contoso")
            .withFilters(Arrays.asList(new KeyValueFilter().withKey("fakeTokenPlaceholder").withLabel("Production")))
            .withRetentionPeriod(3600L).create();
    }
}
