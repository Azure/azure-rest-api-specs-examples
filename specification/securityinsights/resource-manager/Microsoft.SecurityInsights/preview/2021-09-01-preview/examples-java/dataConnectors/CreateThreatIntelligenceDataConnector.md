```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.TIDataConnector;
import com.azure.resourcemanager.securityinsights.models.TIDataConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.TIDataConnectorDataTypesIndicators;
import java.time.OffsetDateTime;

/** Samples for DataConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CreateThreatIntelligenceDataConnector.json
     */
    /**
     * Sample code: Creates or updates an Threat Intelligence Platform data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnThreatIntelligencePlatformDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new TIDataConnector()
                    .withTipLookbackPeriod(OffsetDateTime.parse("2020-01-01T13:00:30.123Z"))
                    .withDataTypes(
                        new TIDataConnectorDataTypes()
                            .withIndicators(new TIDataConnectorDataTypesIndicators().withState(DataTypeState.ENABLED)))
                    .withTenantId("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
