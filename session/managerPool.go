package session

var (
	managers = make(map[string]*Manager)
)

// RegisterSessionManager register new Manager with provider name and json config string.
// use GetSessionManager() to get this manager.
//
// provider name:
// TODO 1. cookie
// TODO 2. file
// TODO 3. memory
// 4. redis providerConfig like "addr:port,poolSize,pwd,dbNum,idleTimeout", e.g. 127.0.0.1:6379,100,pwd,0,30
// TODO 5. mysql
func RegisterSessionManager(sessionKey, providerName string, cf *ManagerConfig) error {
	m, err := NewManager(providerName, cf)
	if err != nil {
		return err
	}

	managers[sessionKey] = m
	return nil
}

func GetSessionManager(sessionKey string) *Manager {
	if _, ok := managers[sessionKey]; !ok {
		return nil
	}
	return managers[sessionKey]
}
