```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresql.fluent.models.ConfigurationListResultInner;

/** Samples for ServerParameters ListUpdateConfigurations. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ConfigurationsUpdateByServer.json
     */
    /**
     * Sample code: ConfigurationList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void configurationList(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverParameters()
            .listUpdateConfigurations("TestGroup", "testserver", new ConfigurationListResultInner(), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.
