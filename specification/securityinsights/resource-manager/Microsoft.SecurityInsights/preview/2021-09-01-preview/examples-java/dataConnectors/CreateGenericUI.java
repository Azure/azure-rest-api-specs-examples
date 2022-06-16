import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.securityinsights.models.Availability;
import com.azure.resourcemanager.securityinsights.models.AvailabilityStatus;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigProperties;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesConnectivityCriteriaItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesDataTypesItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesGraphQueriesItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesInstructionStepsItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiConnectorConfigPropertiesSampleQueriesItem;
import com.azure.resourcemanager.securityinsights.models.CodelessUiDataConnector;
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
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CreateGenericUI.json
     */
    /**
     * Sample code: Creates or updates a GenericUI data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAGenericUIDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) throws IOException {
        manager
            .dataConnectors()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "316ec55e-7138-4d63-ab18-90c8a60fd1c8",
                new CodelessUiDataConnector()
                    .withConnectorUiConfig(
                        new CodelessUiConnectorConfigProperties()
                            .withTitle("Qualys Vulnerability Management (CCP DEMO)")
                            .withPublisher("Qualys")
                            .withDescriptionMarkdown(
                                "The [Qualys Vulnerability Management"
                                    + " (VM)](https://www.qualys.com/apps/vulnerability-management/) data connector"
                                    + " provides the capability to ingest vulnerability host detection data into Azure"
                                    + " Sentinel through the Qualys API. The connector provides visibility into host"
                                    + " detection data from vulerability scans. This connector provides Azure Sentinel"
                                    + " the capability to view dashboards, create custom alerts, and improve"
                                    + " investigation ")
                            .withGraphQueriesTableName("QualysHostDetection_CL")
                            .withGraphQueries(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesGraphQueriesItem()
                                            .withMetricName("Total data received")
                                            .withLegend("{{graphQueriesTableName}}")
                                            .withBaseQuery("{{graphQueriesTableName}}")))
                            .withSampleQueries(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesSampleQueriesItem()
                                            .withDescription("Top 10 Vulerabilities detected")
                                            .withQuery(
                                                "{{graphQueriesTableName}}\n"
                                                    + " | mv-expand todynamic(Detections_s)\n"
                                                    + " | extend Vulnerability = tostring(Detections_s.Results)\n"
                                                    + " | summarize count() by Vulnerability\n"
                                                    + " | top 10 by count_")))
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
                                            .withType(ConnectivityType.IS_CONNECTED_QUERY)
                                            .withValue(
                                                Arrays
                                                    .asList(
                                                        "{{graphQueriesTableName}}\n"
                                                            + "            | summarize LastLogReceived ="
                                                            + " max(TimeGenerated)\n"
                                                            + "            | project IsConnected = LastLogReceived >"
                                                            + " ago(30d)"))))
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
                                                        "read and write permissions on the workspace are required.")
                                                    .withProviderDisplayName("Workspace")
                                                    .withScope(PermissionProviderScope.WORKSPACE)
                                                    .withRequiredPermissions(
                                                        new RequiredPermissions()
                                                            .withWrite(true)
                                                            .withRead(true)
                                                            .withDelete(true)),
                                                new PermissionsResourceProviderItem()
                                                    .withProvider(
                                                        ProviderName
                                                            .MICROSOFT_OPERATIONAL_INSIGHTS_WORKSPACES_SHARED_KEYS)
                                                    .withPermissionsDisplayText(
                                                        "read permissions to shared keys for the workspace are"
                                                            + " required. [See the documentation to learn more about"
                                                            + " workspace"
                                                            + " keys](https://docs.microsoft.com/azure/azure-monitor/platform/agent-windows#obtain-workspace-id-and-key).")
                                                    .withProviderDisplayName("Keys")
                                                    .withScope(PermissionProviderScope.WORKSPACE)
                                                    .withRequiredPermissions(
                                                        new RequiredPermissions().withAction(true))))
                                    .withCustoms(
                                        Arrays
                                            .asList(
                                                new PermissionsCustomsItem()
                                                    .withName("Microsoft.Web/sites permissions")
                                                    .withDescription(
                                                        "Read and write permissions to Azure Functions to create a"
                                                            + " Function App is required. [See the documentation to"
                                                            + " learn more about Azure"
                                                            + " Functions](https://docs.microsoft.com/azure/azure-functions/)."),
                                                new PermissionsCustomsItem()
                                                    .withName("Qualys API Key")
                                                    .withDescription(
                                                        "A Qualys VM API username and password is required. [See the"
                                                            + " documentation to learn more about Qualys VM"
                                                            + " API](https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf)."))))
                            .withInstructionSteps(
                                Arrays
                                    .asList(
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                ">**NOTE:** This connector uses Azure Functions to connect to Qualys VM"
                                                    + " to pull its logs into Azure Sentinel. This might result in"
                                                    + " additional data ingestion costs. Check the [Azure Functions"
                                                    + " pricing"
                                                    + " page](https://azure.microsoft.com/pricing/details/functions/)"
                                                    + " for details."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                ">**(Optional Step)** Securely store workspace and API authorization"
                                                    + " key(s) or token(s) in Azure Key Vault. Azure Key Vault provides"
                                                    + " a secure mechanism to store and retrieve key values. [Follow"
                                                    + " these"
                                                    + " instructions](https://docs.microsoft.com/azure/app-service/app-service-key-vault-references)"
                                                    + " to use Azure Key Vault with an Azure Function App."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                "**STEP 1 - Configuration steps for the Qualys VM API**\n\n"
                                                    + "1. Log into the Qualys Vulnerability Management console with an"
                                                    + " administrator account, select the **Users** tab and the"
                                                    + " **Users** subtab. \n"
                                                    + "2. Click on the **New** drop-down menu and select **Users..**\n"
                                                    + "3. Create a username and password for the API account. \n"
                                                    + "4. In the **User Roles** tab, ensure the account role is set to"
                                                    + " **Manager** and access is allowed to **GUI** and **API**\n"
                                                    + "4. Log out of the administrator account and log into the console"
                                                    + " with the new API credentials for validation, then log out of"
                                                    + " the API account. \n"
                                                    + "5. Log back into the console using an administrator account and"
                                                    + " modify the API accounts User Roles, removing access to **GUI**."
                                                    + " \n"
                                                    + "6. Save all changes."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                "**STEP 2 - Choose ONE from the following two deployment options to"
                                                    + " deploy the connector and the associated Azure Function**\n\n"
                                                    + ">**IMPORTANT:** Before deploying the Qualys VM connector, have"
                                                    + " the Workspace ID and Workspace Primary Key (can be copied from"
                                                    + " the following), as well as the Qualys VM API Authorization"
                                                    + " Key(s), readily available.")
                                            .withInstructions(
                                                Arrays
                                                    .asList(
                                                        new InstructionStepsInstructionsItem()
                                                            .withParameters(
                                                                SerializerFactory
                                                                    .createDefaultManagementSerializerAdapter()
                                                                    .deserialize(
                                                                        "{\"fillWith\":[\"WorkspaceId\"],\"label\":\"Workspace"
                                                                            + " ID\"}",
                                                                        Object.class,
                                                                        SerializerEncoding.JSON))
                                                            .withType(SettingType.COPYABLE_LABEL),
                                                        new InstructionStepsInstructionsItem()
                                                            .withParameters(
                                                                SerializerFactory
                                                                    .createDefaultManagementSerializerAdapter()
                                                                    .deserialize(
                                                                        "{\"fillWith\":[\"PrimaryKey\"],\"label\":\"Primary"
                                                                            + " Key\"}",
                                                                        Object.class,
                                                                        SerializerEncoding.JSON))
                                                            .withType(SettingType.COPYABLE_LABEL))),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("Option 1 - Azure Resource Manager (ARM) Template")
                                            .withDescription(
                                                "Use this method for automated deployment of the Qualys VM connector"
                                                    + " using an ARM Tempate.\n\n"
                                                    + "1. Click the **Deploy to Azure** button below. \n\n"
                                                    + "\t[![Deploy To"
                                                    + " Azure](https://aka.ms/deploytoazurebutton)](https://aka.ms/sentinelqualysvmazuredeploy)\n"
                                                    + "2. Select the preferred **Subscription**, **Resource Group** and"
                                                    + " **Location**. \n"
                                                    + "3. Enter the **Workspace ID**, **Workspace Key**, **API"
                                                    + " Username**, **API Password** , update the **URI**, and any"
                                                    + " additional URI **Filter Parameters** (each filter should be"
                                                    + " separated by an \"&\" symbol, no spaces.) \n"
                                                    + "> - Enter the URI that corresponds to your region. The complete"
                                                    + " list of API Server URLs can be [found"
                                                    + " here](https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf#G4.735348)"
                                                    + " -- There is no need to add a time suffix to the URI, the"
                                                    + " Function App will dynamically append the Time Value to the URI"
                                                    + " in the proper format. \n"
                                                    + " - The default **Time Interval** is set to pull the last five"
                                                    + " (5) minutes of data. If the time interval needs to be modified,"
                                                    + " it is recommended to change the Function App Timer Trigger"
                                                    + " accordingly (in the function.json file, post deployment) to"
                                                    + " prevent overlapping data ingestion. \n"
                                                    + "> - Note: If using Azure Key Vault secrets for any of the values"
                                                    + " above, use the`@Microsoft.KeyVault(SecretUri={Security"
                                                    + " Identifier})`schema in place of the string values. Refer to"
                                                    + " [Key Vault references"
                                                    + " documentation](https://docs.microsoft.com/azure/app-service/app-service-key-vault-references)"
                                                    + " for further details. \n"
                                                    + "4. Mark the checkbox labeled **I agree to the terms and"
                                                    + " conditions stated above**. \n"
                                                    + "5. Click **Purchase** to deploy."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("Option 2 - Manual Deployment of Azure Functions")
                                            .withDescription(
                                                "Use the following step-by-step instructions to deploy the Quayls VM"
                                                    + " connector manually with Azure Functions."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                "**1. Create a Function App**\n\n"
                                                    + "1.  From the Azure Portal, navigate to [Function"
                                                    + " App](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.Web%2Fsites/kind/functionapp),"
                                                    + " and select **+ Add**.\n"
                                                    + "2. In the **Basics** tab, ensure Runtime stack is set to"
                                                    + " **Powershell Core**. \n"
                                                    + "3. In the **Hosting** tab, ensure the **Consumption"
                                                    + " (Serverless)** plan type is selected.\n"
                                                    + "4. Make other preferrable configuration changes, if needed, then"
                                                    + " click **Create**."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                "**2. Import Function App Code**\n\n"
                                                    + "1. In the newly created Function App, select **Functions** on"
                                                    + " the left pane and click **+ New Function**.\n"
                                                    + "2. Select **Timer Trigger**.\n"
                                                    + "3. Enter a unique Function **Name** and leave the default cron"
                                                    + " schedule of every 5 minutes, then click **Create**.\n"
                                                    + "5. Click on **Code + Test** on the left pane. \n"
                                                    + "6. Copy the [Function App"
                                                    + " Code](https://aka.ms/sentinelqualysvmazurefunctioncode) and"
                                                    + " paste into the Function App `run.ps1` editor.\n"
                                                    + "7. Click **Save**."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                "**3. Configure the Function App**\n\n"
                                                    + "1. In the Function App, select the Function App Name and select"
                                                    + " **Configuration**.\n"
                                                    + "2. In the **Application settings** tab, select **+ New"
                                                    + " application setting**.\n"
                                                    + "3. Add each of the following seven (7) application settings"
                                                    + " individually, with their respective string values"
                                                    + " (case-sensitive): \n"
                                                    + "\t\tapiUsername\n"
                                                    + "\t\tapiPassword\n"
                                                    + "\t\tworkspaceID\n"
                                                    + "\t\tworkspaceKey\n"
                                                    + "\t\turi\n"
                                                    + "\t\tfilterParameters\n"
                                                    + "\t\ttimeInterval\n"
                                                    + "> - Enter the URI that corresponds to your region. The complete"
                                                    + " list of API Server URLs can be [found"
                                                    + " here](https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf#G4.735348)."
                                                    + " The `uri` value must follow the following schema: `https://<API"
                                                    + " Server>/api/2.0/fo/asset/host/vm/detection/?action=list&vm_processed_after=`"
                                                    + " -- There is no need to add a time suffix to the URI, the"
                                                    + " Function App will dynamically append the Time Value to the URI"
                                                    + " in the proper format.\n"
                                                    + "> - Add any additional filter parameters, for the"
                                                    + " `filterParameters` variable, that need to be appended to the"
                                                    + " URI. Each parameter should be seperated by an \"&\" symbol and"
                                                    + " should not include any spaces.\n"
                                                    + "> - Set the `timeInterval` (in minutes) to the value of `5` to"
                                                    + " correspond to the Timer Trigger of every `5` minutes. If the"
                                                    + " time interval needs to be modified, it is recommended to change"
                                                    + " the Function App Timer Trigger accordingly to prevent"
                                                    + " overlapping data ingestion.\n"
                                                    + "> - Note: If using Azure Key Vault, use"
                                                    + " the`@Microsoft.KeyVault(SecretUri={Security Identifier})`schema"
                                                    + " in place of the string values. Refer to [Key Vault references"
                                                    + " documentation](https://docs.microsoft.com/azure/app-service/app-service-key-vault-references)"
                                                    + " for further details.\n"
                                                    + "4. Once all application settings have been entered, click"
                                                    + " **Save**."),
                                        new CodelessUiConnectorConfigPropertiesInstructionStepsItem()
                                            .withTitle("")
                                            .withDescription(
                                                "**4. Configure the host.json**.\n\n"
                                                    + "Due to the potentially large amount of Qualys host detection"
                                                    + " data being ingested, it can cause the execution time to surpass"
                                                    + " the default Function App timeout of five (5) minutes. Increase"
                                                    + " the default timeout duration to the maximum of ten (10)"
                                                    + " minutes, under the Consumption Plan, to allow more time for the"
                                                    + " Function App to execute.\n\n"
                                                    + "1. In the Function App, select the Function App Name and select"
                                                    + " the **App Service Editor** blade.\n"
                                                    + "2. Click **Go** to open the editor, then select the"
                                                    + " **host.json** file under the **wwwroot** directory.\n"
                                                    + "3. Add the line `\"functionTimeout\": \"00:10:00\",` above the"
                                                    + " `managedDependancy` line \n"
                                                    + "4. Ensure **SAVED** appears on the top right corner of the"
                                                    + " editor, then exit the editor.\n\n"
                                                    + "> NOTE: If a longer timeout duration is required, consider"
                                                    + " upgrading to an [App Service"
                                                    + " Plan](https://docs.microsoft.com/azure/azure-functions/functions-scale#timeout)")))),
                Context.NONE);
    }
}
