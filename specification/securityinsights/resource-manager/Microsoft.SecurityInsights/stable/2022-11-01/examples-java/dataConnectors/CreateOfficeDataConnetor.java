
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.OfficeDataConnector;
import com.azure.resourcemanager.securityinsights.models.OfficeDataConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.OfficeDataConnectorDataTypesExchange;
import com.azure.resourcemanager.securityinsights.models.OfficeDataConnectorDataTypesSharePoint;
import com.azure.resourcemanager.securityinsights.models.OfficeDataConnectorDataTypesTeams;

/**
 * Samples for DataConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * dataConnectors/CreateOfficeDataConnetor.json
     */
    /**
     * Sample code: Creates or updates an Office365 data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnOffice365DataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().createOrUpdateWithResponse("myRg", "myWorkspace",
            "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            new OfficeDataConnector().withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                .withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8")
                .withDataTypes(new OfficeDataConnectorDataTypes()
                    .withExchange(new OfficeDataConnectorDataTypesExchange().withState(DataTypeState.ENABLED))
                    .withSharePoint(new OfficeDataConnectorDataTypesSharePoint().withState(DataTypeState.ENABLED))
                    .withTeams(new OfficeDataConnectorDataTypesTeams().withState(DataTypeState.ENABLED))),
            com.azure.core.util.Context.NONE);
    }
}
