import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.DatabaseEdition;
import com.azure.resourcemanager.sql.models.ImportRequest;
import com.azure.resourcemanager.sql.models.ServiceObjectiveName;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Databases ImportMethod. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseCreateImportMinStorageKey.json
     */
    /**
     * Sample code: Import bacpac into new database Min with storage key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importBacpacIntoNewDatabaseMinWithStorageKey(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .importMethod(
                "sqlcrudtest-4799",
                "sqlcrudtest-5961",
                new ImportRequest()
                    .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                    .withStorageKey(
                        "sdlfkjdsf+sdlfkjsdlkfsjdfLDKFJSDLKFDFKLjsdfksjdflsdkfD2342309432849328479324/3RSD==")
                    .withStorageUri("https://test.blob.core.windows.net/bacpacs/testbacpac.bacpac")
                    .withAdministratorLogin("dummyLogin")
                    .withAdministratorLoginPassword("Un53cuRE!")
                    .withDatabaseName("TestDbImport")
                    .withEdition(DatabaseEdition.BASIC)
                    .withServiceObjectiveName(ServiceObjectiveName.BASIC)
                    .withMaxSizeBytes("2147483648"),
                Context.NONE);
    }
}
