
import com.azure.resourcemanager.databasewatcher.models.AlertRuleCreationProperties;
import com.azure.resourcemanager.databasewatcher.models.AlertRuleResourceProperties;
import java.time.OffsetDateTime;

/**
 * Samples for AlertRuleResources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/AlertRuleResources_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: AlertRuleResources_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void alertRuleResourcesCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.alertRuleResources().define("testAlert").withExistingWatcher("rgWatcher", "testWatcher")
            .withProperties(new AlertRuleResourceProperties().withAlertRuleResourceId(
                "/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878Be/resourceGroups/rgWatcher/providers/microsoft.insights/scheduledqueryrules/alerts-demo")
                .withCreatedWithProperties(AlertRuleCreationProperties.CREATED_WITH_ACTION_GROUP)
                .withCreationTime(OffsetDateTime.parse("2024-07-25T15:38:47.798Z"))
                .withAlertRuleTemplateId("someTemplateId").withAlertRuleTemplateVersion("1.0"))
            .create();
    }
}
