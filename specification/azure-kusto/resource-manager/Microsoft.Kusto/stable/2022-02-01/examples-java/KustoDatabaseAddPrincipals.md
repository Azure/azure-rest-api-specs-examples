```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.fluent.models.DatabasePrincipalInner;
import com.azure.resourcemanager.kusto.models.DatabasePrincipalListRequest;
import com.azure.resourcemanager.kusto.models.DatabasePrincipalRole;
import com.azure.resourcemanager.kusto.models.DatabasePrincipalType;
import java.util.Arrays;

/** Samples for Databases AddPrincipals. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabaseAddPrincipals.json
     */
    /**
     * Sample code: KustoDatabaseAddPrincipals.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabaseAddPrincipals(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .addPrincipalsWithResponse(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                new DatabasePrincipalListRequest()
                    .withValue(
                        Arrays
                            .asList(
                                new DatabasePrincipalInner()
                                    .withRole(DatabasePrincipalRole.ADMIN)
                                    .withName("Some User")
                                    .withType(DatabasePrincipalType.USER)
                                    .withFqn("aaduser=some_guid")
                                    .withEmail("user@microsoft.com")
                                    .withAppId(""),
                                new DatabasePrincipalInner()
                                    .withRole(DatabasePrincipalRole.VIEWER)
                                    .withName("Kusto")
                                    .withType(DatabasePrincipalType.GROUP)
                                    .withFqn("aadgroup=some_guid")
                                    .withEmail("kusto@microsoft.com")
                                    .withAppId(""),
                                new DatabasePrincipalInner()
                                    .withRole(DatabasePrincipalRole.ADMIN)
                                    .withName("SomeApp")
                                    .withType(DatabasePrincipalType.APP)
                                    .withFqn("aadapp=some_guid_app_id")
                                    .withEmail("")
                                    .withAppId("some_guid_app_id"))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
