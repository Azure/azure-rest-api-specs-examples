Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.PollingFrequency;
import com.azure.resourcemanager.securityinsights.models.TiTaxiiDataConnector;
import com.azure.resourcemanager.securityinsights.models.TiTaxiiDataConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.TiTaxiiDataConnectorDataTypesTaxiiClient;
import java.time.OffsetDateTime;

/** Samples for DataConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CreateThreatIntelligenceTaxiiDataConnector.json
     */
    /**
     * Sample code: Creates or updates a Threat Intelligence Taxii data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAThreatIntelligenceTaxiiDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new TiTaxiiDataConnector()
                    .withEtag("d12423f6-a60b-4ca5-88c0-feb1a182d0f0")
                    .withWorkspaceId("dd124572-4962-4495-9bd2-9dade12314b4")
                    .withFriendlyName("testTaxii")
                    .withTaxiiServer("https://limo.anomali.com/api/v1/taxii2/feeds")
                    .withCollectionId("135")
                    .withUsername("--")
                    .withPassword("--")
                    .withTaxiiLookbackPeriod(OffsetDateTime.parse("2020-01-01T13:00:30.123Z"))
                    .withPollingFrequency(PollingFrequency.ONCE_ADAY)
                    .withDataTypes(
                        new TiTaxiiDataConnectorDataTypes()
                            .withTaxiiClient(
                                new TiTaxiiDataConnectorDataTypesTaxiiClient().withState(DataTypeState.ENABLED)))
                    .withTenantId("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
                Context.NONE);
    }
}
```
