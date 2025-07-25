
import com.azure.resourcemanager.cloudhealth.models.AlertConfiguration;
import com.azure.resourcemanager.cloudhealth.models.AlertSeverity;
import com.azure.resourcemanager.cloudhealth.models.AzureMonitorWorkspaceSignalGroup;
import com.azure.resourcemanager.cloudhealth.models.AzureResourceSignalGroup;
import com.azure.resourcemanager.cloudhealth.models.DependenciesAggregationType;
import com.azure.resourcemanager.cloudhealth.models.DependenciesSignalGroup;
import com.azure.resourcemanager.cloudhealth.models.EntityAlerts;
import com.azure.resourcemanager.cloudhealth.models.EntityCoordinates;
import com.azure.resourcemanager.cloudhealth.models.EntityImpact;
import com.azure.resourcemanager.cloudhealth.models.EntityProperties;
import com.azure.resourcemanager.cloudhealth.models.IconDefinition;
import com.azure.resourcemanager.cloudhealth.models.LogAnalyticsSignalGroup;
import com.azure.resourcemanager.cloudhealth.models.SignalAssignment;
import com.azure.resourcemanager.cloudhealth.models.SignalGroup;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Entities CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Entities_CreateOrUpdate.json
     */
    /**
     * Sample code: Entities_CreateOrUpdate.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesCreateOrUpdate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().define("uszrxbdkxesdrxhmagmzywebgbjj").withExistingHealthmodel("rgopenapi", "myHealthModel")
            .withProperties(new EntityProperties().withDisplayName("My entity")
                .withCanvasPosition(new EntityCoordinates().withX(14.0).withY(13.0))
                .withIcon(new IconDefinition().withIconName("Custom").withCustomData("rcitntvapruccrhtxmkqjphbxunkz"))
                .withHealthObjective(62.0D).withImpact(EntityImpact.STANDARD)
                .withLabels(mapOf("key1376", "fakeTokenPlaceholder"))
                .withSignals(new SignalGroup().withAzureResource(new AzureResourceSignalGroup()
                    .withSignalAssignments(
                        Arrays.asList(new SignalAssignment().withSignalDefinitions(Arrays.asList("sigdef1"))))
                    .withAuthenticationSetting("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX")
                    .withAzureResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"))
                    .withAzureLogAnalytics(new LogAnalyticsSignalGroup()
                        .withSignalAssignments(Arrays.asList(new SignalAssignment().withSignalDefinitions(
                            Arrays.asList("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"))))
                        .withAuthenticationSetting("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX")
                        .withLogAnalyticsWorkspaceResourceId(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"))
                    .withAzureMonitorWorkspace(new AzureMonitorWorkspaceSignalGroup()
                        .withSignalAssignments(Arrays.asList(
                            new SignalAssignment().withSignalDefinitions(Arrays.asList("sigdef2")),
                            new SignalAssignment().withSignalDefinitions(Arrays.asList("sigdef3"))))
                        .withAuthenticationSetting("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX")
                        .withAzureMonitorWorkspaceResourceId(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"))
                    .withDependencies(
                        new DependenciesSignalGroup().withAggregationType(DependenciesAggregationType.WORST_OF)))
                .withAlerts(new EntityAlerts().withUnhealthy(new AlertConfiguration().withSeverity(AlertSeverity.SEV1)
                    .withDescription("Alert description")
                    .withActionGroupIds(Arrays.asList(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup")))
                    .withDegraded(new AlertConfiguration().withSeverity(AlertSeverity.SEV4)
                        .withDescription("Alert description")
                        .withActionGroupIds(Arrays.asList(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup")))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
