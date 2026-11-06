
import com.azure.resourcemanager.monitor.fluent.models.DataCollectionRuleResourceInner;
import com.azure.resourcemanager.monitor.models.ColumnDefinition;
import com.azure.resourcemanager.monitor.models.DataCollectionRuleDataSources;
import com.azure.resourcemanager.monitor.models.DataCollectionRuleDestinations;
import com.azure.resourcemanager.monitor.models.DataCollectionRuleReferences;
import com.azure.resourcemanager.monitor.models.DataFlow;
import com.azure.resourcemanager.monitor.models.KnownColumnDefinitionType;
import com.azure.resourcemanager.monitor.models.KnownDataFlowStreams;
import com.azure.resourcemanager.monitor.models.KnownLogFileTextSettingsRecordStartTimestampFormat;
import com.azure.resourcemanager.monitor.models.KnownLogFilesDataSourceFormat;
import com.azure.resourcemanager.monitor.models.KnownStorageBlobLookupType;
import com.azure.resourcemanager.monitor.models.LogAnalyticsDestination;
import com.azure.resourcemanager.monitor.models.LogFileSettingsText;
import com.azure.resourcemanager.monitor.models.LogFilesDataSource;
import com.azure.resourcemanager.monitor.models.LogFilesDataSourceSettings;
import com.azure.resourcemanager.monitor.models.ReferencesSpecEnrichmentData;
import com.azure.resourcemanager.monitor.models.StorageBlob;
import com.azure.resourcemanager.monitor.models.StreamDeclaration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DataCollectionRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRulesCreateEnrichment.json
     */
    /**
     * Sample code: Create or update data collection rule with enrichment.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        createOrUpdateDataCollectionRuleWithEnrichment(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRules().createWithResponse("myResourceGroup", "myCollectionRule",
            new DataCollectionRuleResourceInner().withLocation("eastus")
                .withDescription("A rule showcasing ingestion time enrichment")
                .withDataCollectionEndpointId(
                    "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myDataCollectionEndpoint")
                .withReferences(new DataCollectionRuleReferences().withEnrichmentData(
                    new ReferencesSpecEnrichmentData().withStorageBlobs(Arrays.asList(new StorageBlob().withResourceId(
                        "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourcegroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myenrichmentstorage")
                        .withBlobUrl("https://myenrichmentstorage.blob.core.windows.net/enrichment")
                        .withLookupType(KnownStorageBlobLookupType.STRING).withName("mytextdatastore")))))
                .withStreamDeclarations(mapOf("Custom-TabularDataABC",
                    new StreamDeclaration().withColumns(Arrays.asList(
                        new ColumnDefinition().withName("TimeGenerated").withType(KnownColumnDefinitionType.DATETIME),
                        new ColumnDefinition().withName("Message").withType(KnownColumnDefinitionType.STRING),
                        new ColumnDefinition().withName("AdditionalContext")
                            .withType(KnownColumnDefinitionType.STRING)))))
                .withDataSources(new DataCollectionRuleDataSources().withLogFiles(Arrays.asList(new LogFilesDataSource()
                    .withStreams(Arrays.asList("Custom-TabularDataABC"))
                    .withFilePatterns(Arrays.asList("C:\\JavaLogs\\*\\*.log"))
                    .withFormat(KnownLogFilesDataSourceFormat.TEXT)
                    .withSettings(new LogFilesDataSourceSettings().withText(new LogFileSettingsText()
                        .withRecordStartTimestampFormat(KnownLogFileTextSettingsRecordStartTimestampFormat.ISO_8601)))
                    .withName("myTabularLogDataSource"))))
                .withDestinations(new DataCollectionRuleDestinations()
                    .withLogAnalytics(Arrays.asList(new LogAnalyticsDestination().withWorkspaceResourceId(
                        "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace")
                        .withName("centralWorkspace"))))
                .withDataFlows(Arrays.asList(new DataFlow()
                    .withStreams(Arrays.asList(KnownDataFlowStreams.fromString("Custom-TabularDataABC")))
                    .withDestinations(Arrays.asList("centralWorkspace"))
                    .withTransformKql(
                        "source | extend LookupData = lookup_string_am('mytextdatastore', Message) | project TimeGenerated, Message, AdditionalContext = LookupData.Message")
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
