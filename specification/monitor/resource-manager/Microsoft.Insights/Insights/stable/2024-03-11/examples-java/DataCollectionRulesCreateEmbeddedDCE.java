
import com.azure.resourcemanager.monitor.fluent.models.DataCollectionRuleResourceInner;
import com.azure.resourcemanager.monitor.models.ColumnDefinition;
import com.azure.resourcemanager.monitor.models.DataCollectionRuleDestinations;
import com.azure.resourcemanager.monitor.models.DataFlow;
import com.azure.resourcemanager.monitor.models.KnownColumnDefinitionType;
import com.azure.resourcemanager.monitor.models.KnownDataCollectionRuleResourceKind;
import com.azure.resourcemanager.monitor.models.KnownDataFlowStreams;
import com.azure.resourcemanager.monitor.models.LogAnalyticsDestination;
import com.azure.resourcemanager.monitor.models.StreamDeclaration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DataCollectionRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRulesCreateEmbeddedDCE.json
     */
    /**
     * Sample code: Create or update data collection rule with embedded ingestion endpoints.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void createOrUpdateDataCollectionRuleWithEmbeddedIngestionEndpoints(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRules().createWithResponse("myResourceGroup", "myCollectionRule",
            new DataCollectionRuleResourceInner().withLocation("eastus")
                .withKind(KnownDataCollectionRuleResourceKind.fromString(" Direct"))
                .withDescription("A Direct Ingestion Rule with builtin ingestion fqdns")
                .withStreamDeclarations(mapOf("Custom-LOGS1_CL",
                    new StreamDeclaration().withColumns(Arrays.asList(
                        new ColumnDefinition().withName("Time").withType(KnownColumnDefinitionType.DATETIME),
                        new ColumnDefinition().withName("Computer").withType(KnownColumnDefinitionType.STRING),
                        new ColumnDefinition().withName("AdditionalContext").withType(KnownColumnDefinitionType.STRING),
                        new ColumnDefinition().withName("CounterName").withType(KnownColumnDefinitionType.STRING),
                        new ColumnDefinition().withName("CounterValue").withType(KnownColumnDefinitionType.REAL)))))
                .withDestinations(new DataCollectionRuleDestinations()
                    .withLogAnalytics(Arrays.asList(new LogAnalyticsDestination().withWorkspaceResourceId(
                        "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace")
                        .withName("centralWorkspace"))))
                .withDataFlows(Arrays.asList(new DataFlow()
                    .withStreams(Arrays.asList(KnownDataFlowStreams.fromString("Custom-LOGS1_CL")))
                    .withDestinations(Arrays.asList("myworkspace"))
                    .withTransformKql(
                        "source | extend jsonContext = parse_json(AdditionalContext) | project TimeGenerated = Time, Computer, AdditionalContext = jsonContext, CounterName=tostring(jsonContext.CounterName), CounterValue=toreal(jsonContext.CounterValue)")
                    .withOutputStream("Custom-LOGS1_CL"))),
            com.azure.core.util.Context.NONE);
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
