import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.Office365ProjectConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.Office365ProjectConnectorDataTypesLogs;
import com.azure.resourcemanager.securityinsights.models.Office365ProjectDataConnector;

/** Samples for DataConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/CreateOffice365ProjectDataConnetor.json
     */
    /**
     * Sample code: Creates or updates an Office365 Project data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnOffice365ProjectDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new Office365ProjectDataConnector()
                    .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                    .withDataTypes(
                        new Office365ProjectConnectorDataTypes()
                            .withLogs(new Office365ProjectConnectorDataTypesLogs().withState(DataTypeState.ENABLED)))
                    .withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
                Context.NONE);
    }
}
