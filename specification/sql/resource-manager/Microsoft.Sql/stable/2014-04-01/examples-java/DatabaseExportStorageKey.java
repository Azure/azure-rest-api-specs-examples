import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.AuthenticationType;
import com.azure.resourcemanager.sql.models.ExportRequest;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Databases Export. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseExportStorageKey.json
     */
    /**
     * Sample code: Export a database into a new bacpac file with storage key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportADatabaseIntoANewBacpacFileWithStorageKey(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .export(
                "sqlcrudtest-4799",
                "sqlcrudtest-5961",
                "testdb",
                new ExportRequest()
                    .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                    .withStorageKey(
                        "sdlfkjdsf+sdlfkjsdlkfsjdfLDKFJSDLKFDFKLjsdfksjdflsdkfD2342309432849328479324/3RSD==")
                    .withStorageUri("https://test.blob.core.windows.net/bacpacs/testbacpac.bacpac")
                    .withAdministratorLogin("dummyLogin")
                    .withAdministratorLoginPassword("Un53cuRE!")
                    .withAuthenticationType(AuthenticationType.SQL),
                Context.NONE);
    }
}
