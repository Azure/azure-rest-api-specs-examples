```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.Dynamics365DataConnector;
import com.azure.resourcemanager.securityinsights.models.Dynamics365DataConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.Dynamics365DataConnectorDataTypesDynamics365CdsActivities;

/** Samples for DataConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CreateDynamics365DataConnetor.json
     */
    /**
     * Sample code: Creates or updates a Dynamics365 data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesADynamics365DataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "c2541efb-c9a6-47fe-9501-87d1017d1512",
                new Dynamics365DataConnector()
                    .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                    .withDataTypes(
                        new Dynamics365DataConnectorDataTypes()
                            .withDynamics365CdsActivities(
                                new Dynamics365DataConnectorDataTypesDynamics365CdsActivities()
                                    .withState(DataTypeState.ENABLED)))
                    .withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
