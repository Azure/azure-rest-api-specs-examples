import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.OfficePowerBIConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.OfficePowerBIConnectorDataTypesLogs;
import com.azure.resourcemanager.securityinsights.models.OfficePowerBIDataConnector;

/** Samples for DataConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/CreateOfficePowerBIDataConnector.json
     */
    /**
     * Sample code: Creates or updates an Office PowerBI data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnOfficePowerBIDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new OfficePowerBIDataConnector()
                    .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                    .withDataTypes(
                        new OfficePowerBIConnectorDataTypes()
                            .withLogs(new OfficePowerBIConnectorDataTypesLogs().withState(DataTypeState.ENABLED)))
                    .withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
                Context.NONE);
    }
}
