
import com.azure.resourcemanager.kusto.fluent.models.DatabasePrincipalInner;
import com.azure.resourcemanager.kusto.models.DatabasePrincipalListRequest;
import com.azure.resourcemanager.kusto.models.DatabasePrincipalRole;
import com.azure.resourcemanager.kusto.models.DatabasePrincipalType;
import java.util.Arrays;

/**
 * Samples for Databases AddPrincipals.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDatabaseAddPrincipals.
     * json
     */
    /**
     * Sample code: KustoDatabaseAddPrincipals.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabaseAddPrincipals(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databases().addPrincipalsWithResponse("kustorptest", "kustoCluster", "KustoDatabase8",
            new DatabasePrincipalListRequest().withValue(Arrays.asList(
                new DatabasePrincipalInner().withRole(DatabasePrincipalRole.ADMIN).withName("Some User")
                    .withType(DatabasePrincipalType.USER).withFqn("aaduser=some_guid").withEmail("user@microsoft.com")
                    .withAppId(""),
                new DatabasePrincipalInner().withRole(DatabasePrincipalRole.VIEWER).withName("Kusto")
                    .withType(DatabasePrincipalType.GROUP).withFqn("aadgroup=some_guid")
                    .withEmail("kusto@microsoft.com").withAppId(""),
                new DatabasePrincipalInner().withRole(DatabasePrincipalRole.ADMIN).withName("SomeApp")
                    .withType(DatabasePrincipalType.APP).withFqn("aadapp=some_guid_app_id").withEmail("")
                    .withAppId("some_guid_app_id"))),
            com.azure.core.util.Context.NONE);
    }
}
