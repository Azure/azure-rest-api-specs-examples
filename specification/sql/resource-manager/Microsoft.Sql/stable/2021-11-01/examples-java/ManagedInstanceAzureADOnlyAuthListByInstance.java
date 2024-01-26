
import com.azure.core.util.Context;

/** Samples for ManagedInstanceAzureADOnlyAuthentications ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceAzureADOnlyAuthListByInstance.json
     */
    /**
     * Sample code: Gets a list of Azure Active Directory only authentication object.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsAListOfAzureActiveDirectoryOnlyAuthenticationObject(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceAzureADOnlyAuthentications()
            .listByInstance("Default-SQL-SouthEastAsia", "managedInstance", Context.NONE);
    }
}
