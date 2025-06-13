
import com.azure.resourcemanager.appconfiguration.models.KeyValueFilter;
import java.util.Arrays;

/**
 * Samples for Snapshots Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/
     * ConfigurationStoresCreateSnapshot.json
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
