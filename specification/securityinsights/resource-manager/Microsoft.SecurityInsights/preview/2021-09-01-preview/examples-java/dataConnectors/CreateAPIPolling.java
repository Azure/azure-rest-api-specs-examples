import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.securityinsights.models.Availability;
import com.azure.resourcemanager.securityinsights.models.AvailabilityStatus;
import com.azure.resourcemanager.securityinsights.models.CodelessApiPollingDataConnector;
import com.azure.resourcemanager.securityinsights.models.CodelessConnectorPollingAuthProperties;
import com.azure.resourcemanager.securityinsights.models.CodelessConnectorPollingConfigProperties;
import com.azure.resourcemanager.securityinsights.models.CodelessConnectorPollingPagingProperties;
import com.azure.resourcemanager.securityinsights.models.CodelessConnectorPollingRequestProperties;
import com.azure.resourcemanager.securityinsights.models.CodelessConnectorPollingResponseProperties;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigProperties;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesConnectivityCriteriaItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesDataTypesItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesGraphQueriesItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesInstructionStepsItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesSampleQueriesItem;
import com.azure.resourcemanager.securityinsights.models.ConnectivityType;
import com.azure.resourcemanager.securityinsights.models.InstructionStepsInstructionsItem;
import com.azure.resourcemanager.securityinsights.models.PermissionProviderScope;
import com.azure.resourcemanager.securityinsights.models.Permissions;
import com.azure.resourcemanager.securityinsights.models.PermissionsCustomsItem;
import com.azure.resourcemanager.securityinsights.models.PermissionsResourceProviderItem;
import com.azure.resourcemanager.securityinsights.models.ProviderName;
import com.azure.resourcemanager.securityinsights.models.RequiredPermissions;
import com.azure.resourcemanager.securityinsights.models.SettingType;
import java.io.IOException;
import java.util.Arrays;

/** Samples for DataConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CreateAPIPolling.json
     */
    /**
     * Sample code: Creates or updates a APIPolling data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAAPIPollingDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) throws IOException {
        manager
            .dataConnectors()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "316ec55e-7138-4d63-ab18-90c8a60fd1c8",
                new CodelessApiPollingDataConnector()
                    .withConnectorUiConfig(
                        new CodelessUiConnectorConfigProperties()
                            .withTitle("GitHub Enterprise Audit Log")
                            .withPublisher("GitHub")
                            .withDescriptionMarkdown(
                                "The GitHub audit log connector provides the capability to ingest GitHub logs into"
                                    + " Azure Sentinel. By connecting GitHub audit logs into Azure Sentinel, you can"
                                    + " view this data in workbooks, use it to create custom alerts, and improve your"
                                    + " investigation process.")
                            .withGraphQueriesTableName("GitHubAuditLogPolling_CL")
                            .withGraphQueries(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesGraphQueriesItem()
                                            .withMetricName("Total events received")
                                            .withLegend("GitHub audit log events")
                                            .withBaseQuery("{{graphQueriesTableName}}")))
                            .withSampleQueries(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesSampleQueriesItem()
                                            .withDescription("All logs")
                                            .withQuery("{{graphQueriesTableName}}\n | take 10 <change>")))
                            .withDataTypes(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesDataTypesItem()
                                            .withName("{{graphQueriesTableName}}")
                                            .withLastDataReceivedQuery(
                                                "{{graphQueriesTableName}}\n"
                                                    + "            | summarize Time = max(TimeGenerated)\n"
                                                    + "            | where isnotempty(Time)")))
                            .withConnectivityCriteria(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesConnectivityCriteriaItem()
                                            .withType(ConnectivityType.fromString("SentinelKindsV2"))
                                            .withValue(Arrays.asList())))
                            .withAvailability(new Availability().withStatus(AvailabilityStatus.ONE).withIsPreview(true))
                            .withPermissions(
                                new Permissions()
                                    .withResourceProvider(
                                        Arrays
                                            .asList(
                                                new PermissionsResourceProviderItem()
                                                    .withProvider(
                                                        ProviderName.MICROSOFT_OPERATIONAL_INSIGHTS_WORKSPACES)
                                                    .withPermissionsDisplayText(
                                                        "read and write permissions are required.")
                                                    .withProviderDisplayName("Workspace")
                                                    .withScope(PermissionProviderScope.WORKSPACE)
                                                    .withRequiredPermissions(
                                                        new RequiredPermissions()
                                                            .withWrite(true)
                                                            .withRead(true)
                                                            .withDelete(true))))
                                    .withCustoms(
                                        Arrays
                                            .asList(
                                                new PermissionsCustomsItem()
                                                    .withName("GitHub API personal token Key")
                                                    .withDescription(
                                                        "You need access to GitHub personal token, the key should have"
                                                            + " 'admin:org' scope"))))
                            .withInstructionSteps(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("Connect GitHub Enterprise Audit Log to Azure Sentinel")
                                            .withDescription(
                                                "Enable GitHub audit Logs. \n"
                                                    + " Follow"
                                                    + " [this](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token)"
                                                    + " to create or find your personal key")
                                            .withInstructions(
                                                Arrays
                                                    .asList(
                                                        new InstructionStepsInstructionsItem()
                                                            .withParameters(
                                                                SerializerFactory
                                                                    .createDefaultManagementSerializerAdapter()
                                                                    .deserialize(
                                                                        "{\"enable\":\"true\",\"userRequestPlaceHoldersInput\":[{\"displayText\":\"Organization"
                                                                            + " Name\",\"placeHolderName\":\"{{placeHolder1}}\",\"placeHolderValue\":\"\",\"requestObjectKey\":\"apiEndpoint\"}]}",
                                                                        Object.class,
                                                                        SerializerEncoding.JSON))
                                                            .withType(SettingType.fromString("APIKey")))))))
                    .withPollingConfig(
                        new CodelessConnectorPollingConfigProperties()
                            .withAuth(
                                new CodelessConnectorPollingAuthProperties()
                                    .withAuthType("APIKey")
                                    .withApiKeyName("Authorization")
                                    .withApiKeyIdentifier("token"))
                            .withRequest(
                                new CodelessConnectorPollingRequestProperties()
                                    .withApiEndpoint("https://api.github.com/organizations/{{placeHolder1}}/audit-log")
                                    .withRateLimitQps(50)
                                    .withQueryWindowInMin(15)
                                    .withHttpMethod("Get")
                                    .withQueryTimeFormat("yyyy-MM-ddTHH:mm:ssZ")
                                    .withRetryCount(2)
                                    .withTimeoutInSeconds(60)
                                    .withHeaders(
                                        SerializerFactory
                                            .createDefaultManagementSerializerAdapter()
                                            .deserialize(
                                                "{\"Accept\":\"application/json\",\"User-Agent\":\"Scuba\"}",
                                                Object.class,
                                                SerializerEncoding.JSON))
                                    .withQueryParameters(
                                        SerializerFactory
                                            .createDefaultManagementSerializerAdapter()
                                            .deserialize(
                                                "{\"phrase\":\"created:{_QueryWindowStartTime}..{_QueryWindowEndTime}\"}",
                                                Object.class,
                                                SerializerEncoding.JSON)))
                            .withPaging(
                                new CodelessConnectorPollingPagingProperties()
                                    .withPagingType("LinkHeader")
                                    .withPageSizeParaName("per_page"))
                            .withResponse(
                                new CodelessConnectorPollingResponseProperties()
                                    .withEventsJsonPaths(Arrays.asList("$")))),
                Context.NONE);
    }
}
