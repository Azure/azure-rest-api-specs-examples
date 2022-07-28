import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ExtensionName;
import com.azure.resourcemanager.sql.models.ImportExtensionRequest;
import com.azure.resourcemanager.sql.models.ImportOperationMode;
import com.azure.resourcemanager.sql.models.StorageKeyType;

/** Samples for Databases CreateImportOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseUpdateImportMinStorageKey.json
     */
    /**
     * Sample code: Import bacpac into an existing database Min with storage key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importBacpacIntoAnExistingDatabaseMinWithStorageKey(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .createImportOperation(
                "sqlcrudtest-4799",
                "sqlcrudtest-5961",
                "testdb",
                ExtensionName.IMPORT,
                new ImportExtensionRequest()
                    .withOperationMode(ImportOperationMode.IMPORT)
                    .withStorageKeyType(StorageKeyType.STORAGE_ACCESS_KEY)
                    .withStorageKey(
                        "sdlfkjdsf+sdlfkjsdlkfsjdfLDKFJSDLKFDFKLjsdfksjdflsdkfD2342309432849328479324/3RSD==")
                    .withStorageUri("https://test.blob.core.windows.net/bacpacs/testbacpac.bacpac")
                    .withAdministratorLogin("dummyLogin")
                    .withAdministratorLoginPassword("Un53cuRE!"),
                Context.NONE);
    }
}
